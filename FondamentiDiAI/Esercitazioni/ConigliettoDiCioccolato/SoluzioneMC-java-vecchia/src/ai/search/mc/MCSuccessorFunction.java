/*
 * Created on May 12, 2005
 *
 */
package ai.search.mc;


import java.util.List;
import java.util.ArrayList;


import aima.search.framework.Successor;
import aima.search.framework.SuccessorFunction;

/**
 * @author fchesani
 *
 */
public class MCSuccessorFunction implements SuccessorFunction {
	
/*****************************************************************************
* CONSTRUCTORS
/*****************************************************************************/
	
	public MCSuccessorFunction() {
	};
	
	
	
/*****************************************************************************
* INTERFACE METHODS
/*****************************************************************************/
	public List getSuccessors(Object state) {
		
		
		List result = new ArrayList();
		
		if (state instanceof MCState) {
			MCState mcState = (MCState) state;
			
	 		// missionars that are on the shore where there is also the boat
	 		int numMissionari;
	 		// cannibals that are on the shore where there is also the boat
	 		int numCannibali;
	 		// depending on the shore, I calculate how many missionars/cannibals are
	 		// on that shore
	 		if (mcState.isPosBoat()) {
	 			numMissionari = mcState.getMissionars();
	 			numCannibali = mcState.getCannibals();
	 		}
	 		else {
	 			numMissionari = mcState.getTotMissionars() - mcState.getMissionars();
	 			numCannibali = mcState.getTotCannibals() - mcState.getCannibals();
	 		}
	 			
		  	if ((numMissionari > 0) && (numCannibali > 0)) {
		  		MCState newState = moveMC(mcState);
		  		if (isAllowed(newState))
		  			result.add(new Successor( MCState.mc,newState));
		  	}
	 		if (numMissionari > 1) {
		  		MCState newState = moveMM(mcState);
		  		if (isAllowed(newState))
		  			result.add(new Successor( MCState.mm,newState));
		  	}
	 		if (numCannibali > 1) {
		  		MCState newState = moveCC(mcState);
		  		if (isAllowed(newState))
		  			result.add(new Successor( MCState.cc,newState));
		  	}
	 		if (numMissionari > 0) {
		  		MCState newState = moveM(mcState);
		  		if (isAllowed(newState))
		  			result.add(new Successor( MCState.m,newState));
		  	}
	 		if (numCannibali > 0) {
		  		MCState newState = moveC(mcState);
		  		if (isAllowed(newState))
		  			result.add(new Successor( MCState.c,newState));
		  	}
		}
		return result;
	}
	
	
	
/*****************************************************************************
* PRIVATE METHODS
/*****************************************************************************/

 	// check if there are not too much cannibals...
 	private boolean isAllowed(MCState mcState) {
 		
 		if (	(mcState.getMissionars() == 0) ||
 				(mcState.getMissionars() == mcState.getTotMissionars()) ||
 				(		(mcState.getMissionars() >= mcState.getCannibals())) &&
 						(((mcState.getTotMissionars()-mcState.getMissionars()) >=
 						(mcState.getTotCannibals()-mcState.getCannibals()))
 				)
		)
 			return true;
 		else {
 			return false;
 		}
 	}
	
	
	
/*****************************************************************************
* Methods for applying different moves
/*****************************************************************************/
	private MCState moveMC(MCState mcState) {
		if (mcState.isPosBoat())
 			return new MCState(mcState.getTotMissionars(),
 															mcState.getTotCannibals(),
															mcState.getMissionars() - 1,
															mcState.getCannibals() - 1,
															! mcState.isPosBoat());
 		else
 			return new MCState(mcState.getTotMissionars(),
															mcState.getTotCannibals(),
															mcState.getMissionars() + 1,
															mcState.getCannibals() + 1,
															! mcState.isPosBoat());
 	}
	
 	private MCState moveMM(MCState mcState) {
 		if (mcState.isPosBoat())
 			return new MCState(mcState.getTotMissionars(),
															mcState.getTotCannibals(),
															mcState.getMissionars() - 2,
															mcState.getCannibals(),
															! mcState.isPosBoat());
 		else
 			return new MCState(mcState.getTotMissionars(),
															mcState.getTotCannibals(),
															mcState.getMissionars() + 2,
															mcState.getCannibals(),
															! mcState.isPosBoat());
 	}

 	private MCState moveCC(MCState mcState) {
 		if (mcState.isPosBoat())
 			return new MCState(mcState.getTotMissionars(),
															mcState.getTotCannibals(),
															mcState.getMissionars(),
															mcState.getCannibals() - 2,
															! mcState.isPosBoat());
 		else
 			return new MCState(mcState.getTotMissionars(),
															mcState.getTotCannibals(),
															mcState.getMissionars(),
															mcState.getCannibals() + 2,
															! mcState.isPosBoat());
 	}
	
 	private MCState moveM(MCState mcState) {
 		if (mcState.isPosBoat())
 			return new MCState(mcState.getTotMissionars(),
															mcState.getTotCannibals(),
															mcState.getMissionars() - 1,
															mcState.getCannibals(),
															! mcState.isPosBoat());
 		else
 			return new MCState(mcState.getTotMissionars(),
															mcState.getTotCannibals(),
															mcState.getMissionars() + 1,
															mcState.getCannibals(),
															! mcState.isPosBoat());
 	}
 	
 	private MCState moveC(MCState mcState) {
 		if (mcState.isPosBoat())
 			return new MCState(mcState.getTotMissionars(),
															mcState.getTotCannibals(),
															mcState.getMissionars(),
															mcState.getCannibals() - 1,
															! mcState.isPosBoat());
 		else
 			return new MCState(mcState.getTotMissionars(),
															mcState.getTotCannibals(),
															mcState.getMissionars(),
															mcState.getCannibals() + 1,
															! mcState.isPosBoat());
 	}
	
}
