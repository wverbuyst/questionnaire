import React from "react";
import ReactDOM from "react-dom";
import App from "./App";
import { createOvermind } from "overmind";
import { Provider } from "overmind-react";
import "./index.css";
import { config } from "./state";

const overmind = createOvermind(config);

const rootElement = document.getElementById("root");

ReactDOM.render(
  <React.StrictMode>
    <Provider value={overmind}>
      <App />
    </Provider>
  </React.StrictMode>,
  rootElement
);
