import * as types from './action_types'
import { moveDrone } from './actions'
import mqtt from 'mqtt'

const setupDroneClient = (dispatch) => {
  const client = mqtt.connect(process.env.REACT_APP_SOCKET_URL) // you add a ws:// url here
  client.subscribe(types.DRONE_MOVE)
  client.on("message", function (topic, payload) {
    switch (topic) {
      case types.DRONE_MOVE:
        dispatch(moveDrone(JSON.parse(new TextDecoder("utf-8").decode(payload))))
        break
      default:
        break
    }
  })

  return client
}

export default setupDroneClient