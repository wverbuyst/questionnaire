import { useCallback } from "react";

import "survey-core/defaultV2.min.css";
import { StylesManager, Model } from "survey-core";
import { Survey } from "survey-react-ui";

import "./App.css";

StylesManager.applyTheme("defaultV2");

const surveyJson = {
  elements: [
    {
      name: "FirstName",
      title: "Enter your first name:",
      type: "text",
    },
    {
      name: "LastName",
      title: "Enter your last name:",
      type: "text",
    },
  ],
};

function App() {
  const survey = new Model(surveyJson);
  survey.focusFirstQuestionAutomatic = false;

  // eslint-disable-next-line @typescript-eslint/no-explicit-any
  const alertResults = useCallback((sender: { [key: string]: any }) => {
    const results = JSON.stringify(sender.data);
    alert(results);
  }, []);

  survey.onComplete.add(alertResults);

  return (
    <div className="display">
      <Survey model={survey} />
    </div>
  );
}

export default App;
