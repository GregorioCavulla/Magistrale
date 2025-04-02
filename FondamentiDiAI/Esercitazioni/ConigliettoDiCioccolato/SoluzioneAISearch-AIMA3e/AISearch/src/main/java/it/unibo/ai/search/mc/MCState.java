package it.unibo.ai.search.mc;


import aima.core.agent.Action;
import aima.core.search.framework.evalfunc.HeuristicFunction;
import aima.core.search.framework.problem.GoalTest;
import aima.core.search.framework.problem.StepCostFunction;




public class MCState 
	implements 	GoalTest,
					StepCostFunction,
					HeuristicFunction {

	// Actions
	static public final String mc = "MC";
	static public final String mm = "MM";
	static public final String cc = "CC";
	static public final String m = "M";
	static public final String c = "C";
	
	/*
	public static Action MC = new DynamicAction("MC");
	public static Action MM = new DynamicAction("MM");
	public static Action CC = new DynamicAction("CC");
	public static Action M = new DynamicAction("M");
	public static Action C = new DynamicAction("C");
	*/
	
	
	
	
/*****************************************************************************
 * Description of the state of the problem
/*****************************************************************************/
	/*
	 * Please note that I have a little far extended representation of the state.
	 * 
	 * Through this representation it is possible to map also problems with
	 * a higher number of Missionars and Cannibals, as well as to experiment
	 * different configurations of the initial state (i.e. not all the people on
	 * the same shore in the beginning) 
	 */
	
	// total number of Missionars at the beginning of the game
	private final int totMissionars;
	
	// number of Missionars at the starting shore
	private int missionars;
	
   // total number of Cannibals at the beginning of the game
	private final int totCannibals;
	
	// number of Cannibals at the starting shore
	private int cannibals;
	
	/*
	 * true means that it is on this shore, false means that is on the other shore
	 */
	private boolean posBoat;
 
 
/*****************************************************************************
 * Methods for accessing the state
/*****************************************************************************/
	/**
	 * @return Returns the cannibals at the starting shore.
	 */
	public int getCannibals() {
		return cannibals;
	}
	/**
	 * @return Returns the missionars at the starting shore.
	 */
	public int getMissionars() {
		return missionars;
	}
	/**
	 * @return Returns true if the boat is at the starting shore.
	 */
	public boolean isPosBoat() {
		return posBoat;
	}
	/**
	 * @return Returns the total number ofCannibals.
	 */
	public int getTotCannibals() {
		return totCannibals;
	}
	/**
	 * @return Returns the total number of Missionars.
	 */
	public int getTotMissionars() {
		return totMissionars;
	}
	
	
	
	
	
	
	
/*****************************************************************************
 * CONSTRUCTORS	
/*****************************************************************************/
	/*
	 * Default constructor.
	 * 
	 * It initializes the problem in the classic configuration (3,3,1).
	 */
	public MCState() {
		this(3, 3, 3, 3, true);
	}

	/*
	 * Generic constructor.
	 * 
	 * Used for generating new states, as well as for creating different 
	 * instances of the problem (with more people, and maybe on different shores)
	 */
	public MCState( int totMissionars,
												int totCannibals,
												int missionarsOnThisShore,
												int cannibalsOnThisShore,
												boolean posBoat) {
		this.totMissionars = totMissionars;
		this.missionars = missionarsOnThisShore;
		this.totCannibals = totCannibals;
		this.cannibals = cannibalsOnThisShore;
		this.posBoat = posBoat;
	}
	
	
	
	
	
	
	
/*****************************************************************************
* Methods for the interface GoalTest
/*****************************************************************************/
	public boolean isGoalState(Object state) {
		if (state instanceof MCState) {
			MCState mcState = (MCState) state;
		   return (		(mcState.missionars == 0) && 
		    				(mcState.cannibals == 0) &&
							(mcState.posBoat == false)
			);
		}
		else
			return false;
	};
 
 
 
 

/*****************************************************************************
* Methods for the interface StepCostFunction
/*****************************************************************************/
	/*
 	public Double calculateStepCost(Object fromState, Object toState, String action) {
 		return new Double(1);
 	};
 	*/
 	
 	
 	public double c(Object fromState, Action a, Object toState) {
 		return 1.0;
 	}
 	
 


/*****************************************************************************
* Methods for the interface HeuristicFunction
/*****************************************************************************/
 	// This is the same proposed by Russell-Norvig
	public double getHeuristicValue(Object state) {
		if (state instanceof MCState) {
 			MCState mcState = (MCState) state;
 			int hVal = mcState.missionars + mcState.cannibals - (mcState.posBoat? 1:0);
 			return hVal;
 		}
 		else return Integer.MAX_VALUE;
	}

	public double h(Object state) {
		return this.getHeuristicValue(state);
	}
 
 
/*****************************************************************************
* Generic Methods
/*****************************************************************************/
 	public int hashCode() {
		return 0;
	}

 	public boolean equals(Object o1) {
 		MCState temp = (MCState) o1; 
 		if (	(this.missionars == temp.missionars) &&
 				(this.cannibals == temp.cannibals) &&
				(this.posBoat == temp.posBoat))
 			return true;
 		else
 			return false;
 		
 	}
	
 	public String toString() {
 		return 	"(M:" + this.missionars + 
					", C:" + this.cannibals +
					", B:" + this.posBoat + " )";
	}
 	
 	
 	

}
