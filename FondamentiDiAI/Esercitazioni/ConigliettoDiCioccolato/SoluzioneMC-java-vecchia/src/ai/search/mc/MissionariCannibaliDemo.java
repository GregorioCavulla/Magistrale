/*
 * Created on May 12, 2005
 *
 */
package ai.search.mc;

import java.io.BufferedReader;
import java.io.InputStreamReader;
import java.util.Iterator;
import java.util.List;
import java.util.Properties;

import aima.search.framework.*;
import aima.search.uninformed.*;
import aima.search.informed.*;

/**
 * @author fchesani
 *
 */
public class MissionariCannibaliDemo {
 
	
	public static void main(String[] args) {
		MCState initState = new MCState();

		try {
			BufferedReader br = new BufferedReader(new InputStreamReader(System.in));
			
			
			Problem problem = new Problem(initState,
					new MCSuccessorFunction(),
					initState);
			
			System.out.println("\nPress enter to execute Breadth First Search (Tree Search)...");
			br.readLine();
			System.out.println("\nBreadth First (Tree Search):");
			Search search = new BreadthFirstSearch(new TreeSearch());
			SearchAgent agent = new SearchAgent(problem, search);
			printActions(agent.getActions());
			printInstrumentation(agent.getInstrumentation());
			
			System.out.println("\nPress enter to execute Breadth First Search (Graph Search)...");
			br.readLine();
			System.out.println("\nBreadth First (Graph Search):");
			search = new BreadthFirstSearch(new GraphSearch());
			agent = new SearchAgent(problem, search);
			printActions(agent.getActions());
			printInstrumentation(agent.getInstrumentation());
			
			
			System.out.println("\nPress enter to execute Depth First Search (Graph Search)...");
			br.readLine();
			System.out.println("\nDepth First (Graph Search):");
			search = new DepthFirstSearch(new GraphSearch());
			agent = new SearchAgent(problem, search);
			printActions(agent.getActions());
			printInstrumentation(agent.getInstrumentation());
			
			
			/*
			System.out.println("Depth First (without repetition checks):");
			search = new DepthFirstSearch(new TreeSearch());
			agent = new SearchAgent(problem, search);
			printActions(agent.getActions());
			printInstrumentation(agent.getInstrumentation());
			*/
			
			/*
			Search search = new DepthLimitedSearch(9);
			SearchAgent agent = new SearchAgent(problem, search);
			printActions(agent.getActions());
			printInstrumentation(agent.getInstrumentation());
			*/
			
			problem = new Problem(initState,
											new MCSuccessorFunction(),
											initState,
											initState,
											initState);
			System.out.println("\nPress enter to execute Greedy Best First Search (Tree Search)...");
			br.readLine();
			System.out.println("\nGreedy (Tree Search):");
			search = new GreedyBestFirstSearch(new TreeSearch());
			agent = new SearchAgent(problem, search);
			printActions(agent.getActions());
			printInstrumentation(agent.getInstrumentation());
			
			
			System.out.println("\nPress enter to execute Hill Climbing Search...");
			br.readLine();
			System.out.println("\nHill Climbing:");
			search = new HillClimbingSearch();
			agent = new SearchAgent(problem, search);
			printActions(agent.getActions());
			printInstrumentation(agent.getInstrumentation());
			
			
			System.out.println("\nPress enter to execute A* Search (Tree Search)...");
			br.readLine();
			System.out.println("\nAStar (Tree Search):");
			search = new AStarSearch(new TreeSearch());
			agent = new SearchAgent(problem, search);
			printActions(agent.getActions());
			printInstrumentation(agent.getInstrumentation());
			
			System.out.println("\nPress enter to execute A* Search (Graph Search)...");
			br.readLine();
			System.out.println("\nAStar (Graph Search):");
			search = new AStarSearch(new GraphSearch());
			agent = new SearchAgent(problem, search);
			printActions(agent.getActions());
			printInstrumentation(agent.getInstrumentation());
			
			
		} catch (Exception e) {
			e.printStackTrace();
		}
	}
	
	
	private static void printInstrumentation(Properties properties) {
		Iterator keys = properties.keySet().iterator();
		while (keys.hasNext()) {
			String key = (String) keys.next();
			String property = properties.getProperty(key);
			System.out.println(key + " : " + property.toString());
		}

	}

	private static void printActions(List actions) {
		for (int i = 0; i < actions.size(); i++) {
			String action = (String) actions.get(i);
			System.out.println(action);
		}
	}
}
