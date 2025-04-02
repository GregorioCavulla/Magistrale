package ai.search.u2;

import java.io.BufferedReader;
import java.io.InputStreamReader;
import java.util.Iterator;
import java.util.List;
import java.util.Properties;

import aima.search.framework.*;
import aima.search.uninformed.*;
import aima.search.informed.*;




public class U2Demo {


	 
		
		public static void main(String[] args) {
			U2State initState = new U2State();

			try {
				BufferedReader br = new BufferedReader(new InputStreamReader(System.in));
				
				
				Problem problem = new Problem(initState,
						new U2SuccessorFunction(),
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
				problem = new Problem(	initState,
						new U2SuccessorFunction(),
						initState);
				System.out.println("\nBreadth First (Graph Search):");
				search = new BreadthFirstSearch(new GraphSearch());
				agent = new SearchAgent(problem, search);
				printActions(agent.getActions());
				printInstrumentation(agent.getInstrumentation());
				
				
				System.out.println("\nPress enter to execute Depth First Search (Graph Search)...");
				br.readLine();
				problem = new Problem(	initState,
						new U2SuccessorFunction(),
						initState);
				System.out.println("\nDepth First (Graph Search):");
				search = new DepthFirstSearch(new GraphSearch());
				agent = new SearchAgent(problem, search);
				printActions(agent.getActions());
				printInstrumentation(agent.getInstrumentation());
				
				
				/**/
				System.out.println("\nPress enter to execute Depth First Search (without repetition checks)...");
				br.readLine();
				problem = new Problem(	initState,
						new U2SuccessorFunction(),
						initState);
				System.out.println("Depth First (without repetition checks):");
				search = new DepthFirstSearch(new TreeSearch());
				agent = new SearchAgent(problem, search);
				printActions(agent.getActions());
				printInstrumentation(agent.getInstrumentation());
				/**/
				
				/**/
				System.out.println("\nPress enter to execute Depth First Search (limited depth)...");
				br.readLine();
				problem = new Problem(	initState,
						new U2SuccessorFunction(),
						initState);
				System.out.println("Depth First (limited depth):");
				search = new DepthLimitedSearch(3);
				agent = new SearchAgent(problem, search);
				printActions(agent.getActions());
				printInstrumentation(agent.getInstrumentation());
				/**/
				
				
				
				/**/
				problem = new Problem(initState,
												new U2SuccessorFunction(),
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
				
				System.out.println("\nEnd Search of the U2 Problem\n");
				
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

