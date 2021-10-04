package revision

import (
	"context"
	"github.com/maxh/gqlgen-todos/nodevalue"
	"github.com/maxh/gqlgen-todos/orm/ent"
	"github.com/maxh/gqlgen-todos/qid"
	"strconv"
)

func AddRevision(next ent.Mutator) ent.Mutator {
	return ent.MutateFunc(func(ctx context.Context, m ent.Mutation) (ent.Value, error) {
		// TODO: assert mutation is in a transaction
		// Mutation.Tx

		if m.Type() == ent.TypeNodeRevision {
			// We don't want to store revisions for records for revisions themselves,
			// otherwise we'll end up in an infinite loop.
			return next.Mutate(ctx, m)
		}

		client := ent.FromContext(ctx)

		// Apply the mutation _before_ saving the revision.
		// (Other hooks may change the node before persistence, and we only want to save
		// the "final" revision from this transaction in the revisions table.)
		v, err := next.Mutate(ctx, m)
		if err != nil {
			// Don't save a revision if the mutation failed.
			return v, err
		}

		id := entityId(m)
		node := mutatedNode(ctx, m, client)
		value := toNodeValue(node)
		_, err = client.NodeRevision.Create().
			SetNodeID(string(id)).
			SetNodeRevision("456"). // TODO set revision number
			SetNodeValue(&value).
			Save(ctx)
		if err != nil {
			return nil, err
		}

		return v, nil
	})
}

func toNodeValue(node *ent.Node) nodevalue.NodeValue {
	fieldMap := nodevalue.FieldMap{}
	for _, f := range node.Fields {
		st, err := strconv.Unquote(f.Value)
		if err != nil {
			// Booleans cannot be unquoted; it's not a problem
			// to fallback on the raw value.
			st = f.Value
		}
		fieldMap[f.Name] = st
	}
	edgeMap := nodevalue.EdgeMap{}
	for _, e := range node.Edges {
		edgeMap[e.Name] = e.IDs
	}
	value := nodevalue.NodeValue{
		Fields: fieldMap,
		Edges:  edgeMap,
	}
	return value
}

func mutatedNode(ctx context.Context, m ent.Mutation, client *ent.Client) *ent.Node {
	// All saved entities must have an ID.
	id := entityId(m)

	// We need to look up the table name because WithFixedNodeType expects it.
	tableName := ent.TablesByEntType[m.Type()]
	noder, err := client.Noder(ctx, id, ent.WithFixedNodeType(tableName))

	// At this point, we expect the node to exist in the transaction context.
	if err != nil {
		panic(err)
	}
	node, err := noder.Node(ctx)
	if err != nil {
		panic(err)
	}

	return node
}

type IDer interface {
	ID() (value qid.ID, exists bool)
}

func entityId(m ent.Mutation) qid.ID {
	rev, ok := m.(IDer)
	if !ok {
		panic("no id method on mutated node")
	}
	id, exists := rev.ID()
	if !exists {
		panic("id does not exist on mutated node")
	}
	return id
}
