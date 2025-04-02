package ai.search.u2;

/*
import java.util.ArrayList;

import aima.basic.XYLocation;
import aima.search.nqueens.NQueensBoard;
*/

//import aima.search.*;
//import java.lang.*;
import java.lang.StringBuffer;

import aima.search.framework.GoalTest;
import aima.search.framework.StepCostFunction;
import aima.search.framework.HeuristicFunction;

public class U2State implements GoalTest,
								StepCostFunction,
								HeuristicFunction
{

	static public final int bonoTime = 1;
	static public final int edgeTime = 2;
	static public final int adamTime = 5;
	static public final int larryTime = 10;
	
	
	/*****************************************************************************
	 * Description of the state of the problem
	/*****************************************************************************/
	/*
	 *  position of bono: 0 means that bono is in the backstage,
	 *  1 means that bono is on the stage
	 */
	private int bonoPos = 0;
	private int edgePos = 0;
	private int adamPos = 0;
	private int larryPos = 0;
	private int lightPos = 0;
	
	private int elapsedTime = 0;
	
	
	/*****************************************************************************
	 * CONSTRUCTORS	
	/*****************************************************************************/
	/*
	 * Default constructor.
	 * 
	 * It initializes the problem in the classic configuration.
	 */
	public U2State() {
		this(0, 0, 0, 0, 0, 0);
	}

	/*
	 * Generic constructor.
	 * 
	 * Used for generating new states, as well as for creating different 
	 * instances of the problem (with more people, and maybe on different shores)
	 */
	public U2State( int bonoPos,
					int edgePos,
					int adamPos,
					int larryPos,
					int lightPos,
					int elapsedTime
					)
	{
		this.bonoPos = bonoPos;
		this.edgePos = edgePos;
		this.adamPos = adamPos;
		this.larryPos = larryPos;
		this.lightPos = lightPos;
		this.elapsedTime = elapsedTime;
	}
		
		
	/*****************************************************************************
	* Methods for accessing the state
	/*****************************************************************************/
	public int getBonoPos() {
		return bonoPos;
	}
	public int getEdgePos() {
		return edgePos;
	}
	public int getAdamPos() {
		return adamPos;
	}
	public int getLarryPos() {
		return larryPos;
	}
	public int getLightPos() {
		return lightPos;
	}
	public int getElapsedTime() {
		return elapsedTime;
	}
	void setAdamPos(int adamPos) {
		this.adamPos = adamPos;
	}
	void setBonoPos(int bonoPos) {
		this.bonoPos = bonoPos;
	}
	void setEdgePos(int edgePos) {
		this.edgePos = edgePos;
	}
	void setLarryPos(int larryPos) {
		this.larryPos = larryPos;
	}
	void setLightPos(int lightPos) {
		this.lightPos = lightPos;
	}
	void setElapsedTime(int elapsedTime) {
		this.elapsedTime = elapsedTime;
	}
	
	/*****************************************************************************
	* Methods for the interface GoalTest
	/*****************************************************************************/
	public boolean isGoalState(Object state)
	{
		if (state instanceof U2State) {
			U2State aState = (U2State) state;
			return ( (aState.getBonoPos() == 1) && 
			    	 (aState.getEdgePos() == 1) &&
			    	 (aState.getAdamPos() == 1) &&
			    	 (aState.getLarryPos() == 1) &&
			    	 (aState.getLightPos() == 1)
			    
			);
		}
		else
			return false;
	};
	 
	 
	 
	 

	/*****************************************************************************
	* Methods for the interface StepCostFunction
	/*****************************************************************************/

 	public Double calculateStepCost(Object fromState, Object toState, String action) {
 		int time1 = 100;
 		int time2 = 100;
 		
 		char[] move = action.toCharArray();
 		switch (move[1]) {
 			case 'b' :  time1 = bonoTime;
 						break;
 			case 'e' :  time1 = edgeTime;
 						break;
 			case 'a' :  time1 = adamTime;
						break;
 			case 'l' :  time1 = larryTime;
						break;
			default  :  time1 = 0;
 		}
 		switch (move[2]) {
			case 'b' :  time2 = bonoTime;
						break;
			case 'e' :  time2 = edgeTime;
						break;
			case 'a' :  time2 = adamTime;
						break;
			case 'l' :  time2 = larryTime;
						break;
			default  :  time2 = 0;
		}
 		if (time1 > time2)
 			return new Double(time1);
 		else
 			return new Double(time2);
 	};

	 


	/*****************************************************************************
	* Methods for the interface HeuristicFunction
	/*****************************************************************************/
	public double getHeuristicValue(Object state) 
	{
		StringBuffer sb = new StringBuffer("");
		U2State u2State = (U2State) state;
		
		if (u2State.getLarryPos() == 0) sb.append("l");
		if (u2State.getAdamPos() == 0) sb.append("a");
		if (u2State.getEdgePos() == 0) sb.append("e");
		if (u2State.getBonoPos() == 0) sb.append("b");
		
		String action1 = "fn";
		String action2 = "fn";
		if (sb.length() >= 3) {
			action1 = action1 + sb.charAt(0);
			action2 = action2 + sb.charAt(2);
		}
		else if (sb.length() >= 1) {
			action1 = action1 + sb.charAt(0);
			action2 = action2 + 'n';
		}
		else {
			action1 = action1 + "n";
			action2 = action2 + "n";
		}
		//System.out.print(sb.toString() + "  ");
		//System.out.println("Heuristic: action: " + action+ "\n");
		return 	this.calculateStepCost(null, null, action1).intValue() +
				this.calculateStepCost(null, null, action2).intValue();
	}

	 
	 
	 
	/*****************************************************************************
	* Generic Methods
	/*****************************************************************************/
	/*
 	public int hashCode() {
		return 0;
	}
	*/

	public boolean equals(Object o1) {
		U2State temp = (U2State) o1; 
		if (
			(this.getBonoPos() == temp.getBonoPos()) &&
			(this.getEdgePos() == temp.getEdgePos()) &&
			(this.getAdamPos() == temp.getAdamPos()) &&
			(this.getLarryPos() == temp.getLarryPos()) &&
			(this.getLightPos() == temp.getLightPos())
		)
			return true;
		else
			return false;
	}
		
	public String toString() {
		return 	"(Bono:" + this.getBonoPos() + 
				", Edge:" + this.getEdgePos() +
				", Adam:" + this.getAdamPos() +
				", Larry:" + this.getLarryPos() +
				", Light:" + this.getLightPos() + " )";
	}

	 	
	public boolean isAllowed() {
		if (this.elapsedTime <= 17)
			return true;
		else
			return false;
	}
	 	
	 	

}


	
