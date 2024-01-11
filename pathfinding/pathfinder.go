package pathfinding

import (
	"strings"

	"github.com/GuardianOfWynn/eco-simulator/territory"
	"github.com/victorbetoni/go-streams/streams"
)

type Pathfinder struct {
	From       *territory.Territory
	Target     *territory.Territory
	Claim      territory.Claim
	RouteStyle territory.RouteStyle
}

func (p *Pathfinder) Route() []territory.Territory {

	routes := [][]territory.Territory{}

	for _, conn := range p.From.Connections {
		connTerr := p.Claim.GetTerritory(conn)
		pfinder := Pathfinder{
			From:       connTerr,
			Target:     p.Target,
			Claim:      p.Claim,
			RouteStyle: p.RouteStyle,
		}
		result, vis := pfinder.expand([]territory.Territory{})
		if result {
			routes = append(routes, vis)
		}
	}

	if len(routes) == 0 {
		return []territory.Territory{}
	}

	currentRoute := routes[0]

	for _, route := range routes {
		if len(route) < len(currentRoute) {
			currentRoute = route
		}
	}

	return currentRoute
}

func (p *Pathfinder) expand(visited []territory.Territory) (bool, []territory.Territory) {
	if p.From == nil {
		return false, visited
	}

	visited = append(visited, *p.From)
	for _, conn := range p.From.Connections {
		if strings.ToLower(conn) == strings.ToLower(p.Target.Name) {
			return true, visited
		}
	}
	for _, conn := range p.From.Connections {
		if streams.StreamOf[territory.Territory](visited...).AnyMatch(func(e territory.Territory) bool {
			return strings.ToLower(e.Name) == strings.ToLower(conn)
		}) {
			continue
		}
		connTerr := p.Claim.GetTerritory(conn)
		pfinder := Pathfinder{
			From:       connTerr,
			Target:     p.Target,
			Claim:      p.Claim,
			RouteStyle: p.RouteStyle,
		}
		result, vis := pfinder.expand(visited)
		if result {
			return result, vis
		}
	}

	visited = visited[:len(visited)-1]
	return false, []territory.Territory{}
}

/*

func (p *PathFinder) expand(from, target territory.Territory, visited territory.Territory[]) (bool, territory.Territory[]) {
    if(from === undefined) {
        return {first: false, second: visited}
    }
    visited.push(from);
    let upNode = getNodeAt(from.x, from.y-1, tree);
    let downNode = getNodeAt(from.x, from.y+1, tree);
    let leftNode = getNodeAt(from.x-1, from.y, tree);
    let rightNode = getNodeAt(from.x+1, from.y, tree);

    let connectorType = from.type as ConnectorType;

    if((upNode !== undefined && upNode.id === target.id && isDirectionAllowed(connectorType, Direction.UP))
        || (downNode !== undefined && downNode.id === target.id && isDirectionAllowed(connectorType, Direction.DOWN))
        || (leftNode !== undefined && leftNode.id === target.id && isDirectionAllowed(connectorType, Direction.LEFT))
        || (rightNode !== undefined && rightNode.id === target.id && isDirectionAllowed(connectorType, Direction.RIGHT))) {
            return {first: true, second: visited};
    }

    let upConnector = getConnectorAt(from.x, from.y-1, connectors);
    let downConnector = getConnectorAt(from.x, from.y+1, connectors);
    let leftConnector = getConnectorAt(from.x-1, from.y, connectors);
    let rightConnector = getConnectorAt(from.x+1, from.y, connectors);

    // Direction is relative to the connector, thats why we use Direction.DOWN instead of Direction.UP and etc
    if(upConnector !== undefined && isDirectionAllowed(upConnector.type as ConnectorType, Direction.DOWN) && !wasVisited(upConnector, visited)) {
        let result = expand(upConnector!, target, connectors, tree, visited);
        if(result.first) {
            return result;
        }
    }

    if(rightConnector !== undefined && isDirectionAllowed(rightConnector.type as ConnectorType, Direction.LEFT) && !wasVisited(rightConnector, visited)){
        let result = expand(rightConnector!, target, connectors, tree, visited);
        if(result.first) {
            return result;
        }
    }

    if(leftConnector !== undefined && isDirectionAllowed(leftConnector.type as ConnectorType, Direction.RIGHT) && !wasVisited(leftConnector, visited)) {
        let result = expand(leftConnector!, target, connectors, tree, visited);
        if(result.first) {
            return result;
        }
    }


    if(downConnector !== undefined && isDirectionAllowed(downConnector.type as ConnectorType, Direction.UP) && !wasVisited(downConnector, visited)) {
        let result = expand(downConnector!, target, connectors, tree, visited);
        if(result.first) {
            return result;
        }
    }

    // If the current connector couldnt reach the target node, remove it from visited array.
    visited.splice(visited.length - 1, 1);

    return {first: false, second: []}

}

export function makeTree(selectedNodes: AbilityNode[], tree: AbilityNode[], connectors: AbilityNodeConnector[]): AbilityTree {
    let paths: AbilityNodeConnector[] = [];
    selectedNodes.forEach(node => {
        if(node.links === null) {
            return;
        }
        node.links.map(x => getAbilityNode(x, tree)).forEach(target => {
            if(selectedNodes.includes(target)) {
                paths.push(...findPath(node, target, connectors, tree));
            }
        })
    });
    return new AbilityTree(tree, selectedNodes, paths, []);
}
*/
