package ai.search.u2;


import java.util.List;
import java.util.ArrayList;


import aima.search.framework.Successor;
import aima.search.framework.SuccessorFunction;


/**
 * @author fchesani
 *
 */
public class U2SuccessorFunction implements SuccessorFunction {
	
	
	/*****************************************************************************
	* CONSTRUCTORS
	/*****************************************************************************/	
	public U2SuccessorFunction() {
	};
		
		
		
	/*****************************************************************************
	* INTERFACE METHODS
	/*****************************************************************************/
	public List getSuccessors(Object state) {
			
		List result = new ArrayList();
			
		if (state instanceof U2State) {
			U2State u2State = (U2State) state;
			
			if (u2State.getLightPos() == 1) //the light is on the right side
			{
				if (u2State.getBonoPos() == 1) { // Bono goes back
					U2State newState = move("bbn", u2State);
					if (newState.isAllowed())
						result.add(new Successor( "bbn",newState));
				}
				if (u2State.getEdgePos() == 1) { // Edge goes back
					U2State newState = move("ben", u2State);
					if (newState.isAllowed())
						result.add(new Successor( "ben",newState));
				}
				if (u2State.getAdamPos() == 1) { // Adam goes back
					U2State newState = move("ban", u2State);
					if (newState.isAllowed())
						result.add(new Successor( "ban",newState));
				}
				if (u2State.getLarryPos() == 1) { // Larry goes back
					U2State newState = move("bln", u2State);
					if (newState.isAllowed())
						result.add(new Successor( "bln",newState));
				}
			}
			
			else // the light is on the initial side
			
			{
				if (u2State.getLarryPos() == 0) { // Larry goes forward...
					if (u2State.getAdamPos() == 0) { // ...with Adam
						U2State newState = move("fla", u2State);
						if (newState.isAllowed())
							result.add(new Successor( "fla",newState));
					}
					if (u2State.getEdgePos() == 0) { // ... with Edge
						U2State newState = move("fle", u2State);
						if (newState.isAllowed())
							result.add(new Successor( "fle",newState));
					}
					if (u2State.getBonoPos() == 0) { // ...with Bono
						U2State newState = move("flb", u2State);
						if (newState.isAllowed())
							result.add(new Successor( "flb",newState));
					}
				}
				if (u2State.getAdamPos() == 0) { // Adam goes forward...
					if (u2State.getEdgePos() == 0) { // ... with Edge
						U2State newState = move("fae", u2State);
						if (newState.isAllowed())
							result.add(new Successor( "fae",newState));
					}
					if (u2State.getBonoPos() == 0) { // ...with Bono
						U2State newState = move("fab", u2State);
						if (newState.isAllowed())
							result.add(new Successor( "fab",newState));
					}
				}
				if (u2State.getEdgePos() == 0) { // Edge goes forward...
					if (u2State.getBonoPos() == 0) { // ...with Bono
						U2State newState = move("feb", u2State);
						if (newState.isAllowed())
							result.add(new Successor( "feb",newState));
					}
				}
			}
		}
			
				 	
		return result;
	}
		
		
		
	
	
		
		
		
	/*****************************************************************************
	* Methods for applying different moves
	/*****************************************************************************/
	private U2State move(String actionString, U2State current)
	{ 
		U2State newState = new U2State(
									current.getBonoPos(),
									current.getEdgePos(),
									current.getAdamPos(),
									current.getLarryPos(),
									current.getLightPos(),
									current.getElapsedTime());
		char[] action = actionString.toCharArray();
		
		if (action[0] == 'f') {
			newState.setLightPos(1);	
			switch (action[1]) {
				case 'b' : newState.setBonoPos(1); break;
				case 'e' : newState.setEdgePos(1); break;
				case 'a' : newState.setAdamPos(1); break;
				case 'l' : newState.setLarryPos(1); break;
			}
			switch (action[2]) {
				case 'b' : newState.setBonoPos(1); break;
				case 'e' : newState.setEdgePos(1); break;
				case 'a' : newState.setAdamPos(1); break;
				case 'l' : newState.setLarryPos(1); break;
			}
		}
		else {
			newState.setLightPos(0);
			switch (action[1]) {
				case 'b' : newState.setBonoPos(0); break;
				case 'e' : newState.setEdgePos(0); break;
				case 'a' : newState.setAdamPos(0); break;
				case 'l' : newState.setLarryPos(0); break;
			}
		}
		
		int cost = newState.calculateStepCost(null, null, actionString).intValue();
		newState.setElapsedTime(current.getElapsedTime()+cost);
		return newState;	
	}


	
	
	
	
	
}
