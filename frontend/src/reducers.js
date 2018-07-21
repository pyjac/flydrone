import { combineReducers } from "redux"
import * as types from './action_types'

const drones = (state = [], action) => {
  switch (action.type) {
    case types.DRONE_MOVE:
      let currentState = state;
      let droneIndex = state.findIndex(d => d.id === action.drone.id)
      if (droneIndex === -1) {
        currentState = mergeDroneUpdate([{ active: true, ...action.drone }], state)
        return currentState
      }
      let drone = currentState[droneIndex]
      drone = {...drone, ...action.drone}
      return [
        ...currentState.slice(0, droneIndex),
        drone,
        ...currentState.slice(droneIndex + 1)
      ]

    case types.MARK_DRONE_INACTIVE:
      let droneIndexId = state.findIndex(d => d.id === action.droneId)
      let markedDrone = state[droneIndexId]
      markedDrone = {...markedDrone, active: false}
      return [
        ...state.slice(0, droneIndexId),
        markedDrone,
        ...state.slice(droneIndexId + 1)
      ]
    default:
      return state
  }
}

const mergeDroneUpdate = (currentState, newState) =>{
  let concatDrones = [];
  for (const newDroneIndex in newState){
    for (const oldDroneIndex in currentState)
        if (newState[newDroneIndex].id !== currentState[oldDroneIndex].id) {
            concatDrones.push(newState[newDroneIndex]);
            break;
        }
  }
  return concatDrones.concat(currentState);
}

const reducers = combineReducers({
  drones,
});

export default reducers;