import { createClient } from "@urql/core";
import { Provider } from "@urql/preact";
import { h, render } from "preact";
import { App } from "./App";

const graphqlClient = createClient({
  url: "http://localhost:4015/graphql",
});

render(
  <Provider value={graphqlClient}>
    <App />
  </Provider>,
  document.getElementById("app")!
);
