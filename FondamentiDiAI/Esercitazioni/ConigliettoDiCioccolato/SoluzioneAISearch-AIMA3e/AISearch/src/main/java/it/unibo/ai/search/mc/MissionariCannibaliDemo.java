/*
 * Created on May 12, 2005
 *
 */
package it.unibo.ai.search.mc;

import java.io.BufferedReader;
import java.io.InputStreamReader;
import java.util.Iterator;
import java.util.List;
import java.util.Properties;

//import aima.core.search.framework.*;
//import aima.core.search.uninformed.*;
//import aima.core.search.informed.*;

import aima.core.agent.Action;

import aima.core.search.framework.SearchAgent;
import aima.core.search.framework.problem.Problem;
import aima.core.search.framework.qsearch.GraphSearch;
import aima.core.search.framework.qsearch.TreeSearch;
import aima.core.search.uninformed.BreadthFirstSearch;
import aima.core.search.uninformed.DepthFirstSearch;
import aima.core.search.uninformed.DepthLimitedSearch;
import aima.core.search.informed.AStarSearch;
import aima.core.search.informed.GreedyBestFirstSearch;
import aima.core.search.local.HillClimbingSearch;

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
					new MCActionsFunction(),
					new MCResultFunction(),
					initState,
					initState);
			
			SearchAgent agent;
			
			{
				System.out.println("\nPress enter to execute Breadth First Search (Tree Search)...");
				br.readLine();
				System.out.println("\nBreadth First (Tree Search):");
				BreadthFirstSearch search = new BreadthFirstSearch(new TreeSearch());
				agent = new SearchAgent(problem, search);
				printActions(agent.getActions());
				printInstrumentation(agent.getInstrumentation());
			}
			
			{
				System.out.println("\nPress enter to execute Breadth First Search (Graph Search)...");
				br.readLine();
				System.out.println("\nBreadth First (Graph Search):");
				BreadthFirstSearch search = new BreadthFirstSearch(new GraphSearch());
				agent = new SearchAgent(problem, search);
				printActions(agent.getActions());
				printInstrumentation(agent.getInstrumentation());
			}
			
			{
				System.out.println("\nPress enter to execute Depth First Search (Graph Search)...");
				br.readLine();
				System.out.println("\nDepth First (Graph Search):");
				DepthFirstSearch search = new DepthFirstSearch(new GraphSearch());
				agent = new SearchAgent(problem, search);
				printActions(agent.getActions());
				printInstrumentation(agent.getInstrumentation());
			}
			
			{
				System.out.println("\nPress enter to execute Depth First (without repetition checks):");
				br.readLine();
				System.out.println("\nDepth First (without repetition checks):");
				DepthFirstSearch search = new DepthFirstSearch(new TreeSearch());
				agent = new SearchAgent(problem, search);
				printActions(agent.getActions());
				printInstrumentation(agent.getInstrumentation());
			}
			
			{
				System.out.println("\nPress enter to execute Depth Limited Search...");
				br.readLine();
				System.out.println("\nDepth Limited:");
				DepthLimitedSearch search = new DepthLimitedSearch(9);
				agent = new SearchAgent(problem, search);
				printActions(agent.getActions());
				printInstrumentation(agent.getInstrumentation());
			}
			
			
			{
				System.out.println("\nPress enter to execute Greedy Best First Search (Tree Search)...");
				br.readLine();
				System.out.println("\nGreedy (Tree Search):");
				GreedyBestFirstSearch search = new GreedyBestFirstSearch(new TreeSearch(),initState);
				agent = new SearchAgent(problem, search);
				printActions(agent.getActions());
				printInstrumentation(agent.getInstrumentation());
			}
			
			{
				System.out.println("\nPress enter to execute Hill Climbing Search...");
				br.readLine();
				System.out.println("\nHill Climbing:");
				HillClimbingSearch search = new HillClimbingSearch(initState);
				agent = new SearchAgent(problem, search);
				printActions(agent.getActions());
				printInstrumentation(agent.getInstrumentation());
			}
			
			{
				System.out.println("\nPress enter to execute A* Search (Tree Search)...");
				br.readLine();
				System.out.println("\nAStar (Tree Search):");
				AStarSearch search = new AStarSearch(new TreeSearch(), initState);
				agent = new SearchAgent(problem, search);
				printActions(agent.getActions());
				printInstrumentation(agent.getInstrumentation());
			}
			
			{
				System.out.println("\nPress enter to execute A* Search (Graph Search)...");
				br.readLine();
				System.out.println("\nAStar (Graph Search):");
				AStarSearch search = new AStarSearch(new GraphSearch(), initState);
				agent = new SearchAgent(problem, search);
				printActions(agent.getActions());
				printInstrumentation(agent.getInstrumentation());
			}
			
			
			
		} catch (Exception e) {
			e.printStackTrace();
		}
	}
	
	
	private static void printInstrumentation(Properties properties) {
		Iterator<Object> keys = properties.keySet().iterator();
		while (keys.hasNext()) {
			String key = (String) keys.next();
			String property = properties.getProperty(key);
			System.out.println(key + " : " + property.toString());
		}

	}

	private static void printActions(List<Action> actions) {
		for (int i = 0; i < actions.size(); i++) {
			String action = actions.get(i).toString();
			System.out.println(action);
		}
	}
}
