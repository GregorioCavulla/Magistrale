package it.unibo.ai.search.mc;

import aima.core.agent.Action;
import aima.core.search.framework.problem.ResultFunction;


public class MCResultFunction implements ResultFunction {

	
	public Object result(Object o, Action a) {
		
		MCState mcState = (MCState) o;
		if (MCActionsFunction.MC.equals(a))
			return moveMC(mcState);
		if (MCActionsFunction.MM.equals(a))
			return moveMM(mcState);
		if (MCActionsFunction.CC.equals(a))
			return moveCC(mcState);
		if (MCActionsFunction.M.equals(a))
			return moveM(mcState);
		if (MCActionsFunction.C.equals(a))
			return moveC(mcState);
		
		//noOp
		return mcState;
	}
	


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
