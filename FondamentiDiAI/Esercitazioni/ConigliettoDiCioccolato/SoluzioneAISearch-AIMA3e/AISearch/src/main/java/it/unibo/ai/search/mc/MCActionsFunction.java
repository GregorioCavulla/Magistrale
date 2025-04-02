package it.unibo.ai.search.mc;

import java.util.Set;
import java.util.LinkedHashSet;


import aima.core.agent.Action;
import aima.core.agent.impl.DynamicAction;
import aima.core.search.framework.problem.ActionsFunction;



public class MCActionsFunction implements ActionsFunction {

	public static Action MC = new DynamicAction("MC");
	public static Action MM = new DynamicAction("MM");
	public static Action CC = new DynamicAction("CC");
	public static Action M = new DynamicAction("M");
	public static Action C = new DynamicAction("C");
	
	
	
//	@Override
	public Set<Action> actions(Object o) {
		MCState mcState = (MCState) o;
		Set<Action> result = new LinkedHashSet<Action>();

		if (canMoveMC(mcState))
			result.add(MCActionsFunction.MC);
		if (canMoveMM(mcState))
			result.add(MCActionsFunction.MM);
		if (canMoveCC(mcState))
			result.add(MCActionsFunction.CC);
		if (canMoveM(mcState))
			result.add(MCActionsFunction.M);
		if (canMoveC(mcState))
			result.add(MCActionsFunction.C);
		
		return result;
	}
	
	
	private boolean canMoveMC(MCState mcState) {
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
	 	
	 	// assunzione: lo stato corrente e' allowed
	 	// posso fare MC se e solo se: 
	 	if ((numMissionari >= 1) && (numCannibali >= 1) &&
	 		(numMissionari>=numCannibali) &&
	 		(mcState.getTotMissionars()-numMissionari+1 >= mcState.getTotCannibals()-numCannibali+1)
	 		)
	 		return true;
	 	else
	 		return false;
	}

	
	private boolean canMoveMM(MCState mcState) {
	 	int numMissionari;
	 	int numCannibali;
	 	
	 	if (mcState.isPosBoat()) {
	 		numMissionari = mcState.getMissionars();
	 		numCannibali = mcState.getCannibals();
	 	}
	 	else {
	 		numMissionari = mcState.getTotMissionars() - mcState.getMissionars();
	 		numCannibali = mcState.getTotCannibals() - mcState.getCannibals();
	 	}
	 	if ( 	(numMissionari >= 2) && 
	 			(numMissionari==2 || numMissionari-2 >= numCannibali) &&
	 			(mcState.getTotMissionars()-numMissionari+2 >= mcState.getTotCannibals()-numCannibali)
	 		)
	 		return true;
	 	else
	 		return false;
	}
	

	// quando muovo i cannibali, devo controllare che la sponda di arrivo non diventi sguarnita
	private boolean canMoveCC(MCState mcState) {
	 	int numMissionari;
	 	int numCannibali;
	 	
	 	if (mcState.isPosBoat()) {
	 		numMissionari = mcState.getMissionars();
	 		numCannibali = mcState.getCannibals();
	 	}
	 	else {
	 		numMissionari = mcState.getTotMissionars() - mcState.getMissionars();
	 		numCannibali = mcState.getTotCannibals() - mcState.getCannibals();
	 	}
	 	if ( 	(numCannibali >= 2) && 
	 			(mcState.getTotMissionars() - numMissionari == 0 ||
	 					mcState.getTotMissionars() - numMissionari >= (mcState.getTotCannibals()- numCannibali + 2)
	 			)
	 		)
	 		return true;
	 	else
	 		return false;
	}
	
	
	// quando muovo un missionario, devo controllare che la sponda di partenza non rimanga sguarnita 
	private boolean canMoveM(MCState mcState) {
	 	int numMissionari;
	 	int numCannibali;
	 	
	 	if (mcState.isPosBoat()) {
	 		numMissionari = mcState.getMissionars();
	 		numCannibali = mcState.getCannibals();
	 	}
	 	else {
	 		numMissionari = mcState.getTotMissionars() - mcState.getMissionars();
	 		numCannibali = mcState.getTotCannibals() - mcState.getCannibals();
	 	}
	 	if ( 	(numMissionari >= 1) && 
	 			(numMissionari == 1 || numMissionari-1 >= numCannibali) &&
	 			(mcState.getTotMissionars()-numMissionari+1 >= mcState.getTotCannibals()-numCannibali)
	 		)
	 		return true;
	 	else
	 		return false;
	}
	
	
	// quando muovo un cannibale, devo controllare che la sponda di arrivo non diventi sguarnita
	private boolean canMoveC(MCState mcState) {
	 	int numMissionari;
	 	int numCannibali;
	 	
	 	if (mcState.isPosBoat()) {
	 		numMissionari = mcState.getMissionars();
	 		numCannibali = mcState.getCannibals();
	 	}
	 	else {
	 		numMissionari = mcState.getTotMissionars() - mcState.getMissionars();
	 		numCannibali = mcState.getTotCannibals() - mcState.getCannibals();
	 	}
	 	if ( 	(numCannibali >= 1) && 
	 			(mcState.getTotMissionars() - numMissionari == 0 ||
	 			 mcState.getTotMissionars() - numMissionari >= mcState.getTotCannibals()- numCannibali + 1
	 			)
	 		)
	 		return true;
	 	else
	 		return false;
	}
}
