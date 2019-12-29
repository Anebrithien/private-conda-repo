import { UserApi } from "@/features/user";
import { routerMiddleware } from "connected-react-router";
import { createBrowserHistory } from "history";
import { applyMiddleware, createStore, Store } from "redux";
import { composeWithDevTools } from "redux-devtools-extension";
import thunk from "redux-thunk";

import createReducer from "./rootReducer";

export const history = createBrowserHistory();

function configureStore() {
  const middleware = composeWithDevTools(
    applyMiddleware(thunk, routerMiddleware(history))
  );

  const store = createStore(createReducer(history), middleware);
  initializeStore(store);

  return store;
}

async function initializeStore(store: Store) {
  const dispatch = (action: any) => store.dispatch(action);

  await dispatch(UserApi.loadUser());
}

export default configureStore();