import React from "react";
import ReactDOM from "react-dom/client";
import App from "./App";
import { createOvermind } from "overmind";
import { Provider } from "overmind-react";
import "./index.css";
import { config } from "./state";

const overmind = createOvermind(config);

ReactDOM.createRoot(document.getElementById("root") as HTMLElement).render(
  <React.StrictMode>
    <Provider value={overmind}>
      <App />
    </Provider>
  </React.StrictMode>
);
