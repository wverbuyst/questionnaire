// import { useCallback } from "react";

import "survey-core/defaultV2.min.css";
import { StylesManager, Model } from "survey-core";
import { Survey } from "survey-react-ui";

import "./App.css";
import { useAppState } from "./state";

StylesManager.applyTheme("defaultV2");

function App() {
  const surveyJson = useAppState().questionnaire;
  const survey = surveyJson && new Model(surveyJson);
  // if (survey) {
  //   survey.focusFirstQuestionAutomatic = false;
  // }

  // // eslint-disable-next-line @typescript-eslint/no-explicit-any
  // const results = useCallback((sender: { [key: string]: any }) => {
  //   return JSON.stringify(sender.data);
  // }, []);

  // if (survey) {
  //   survey.onComplete.add(console.log(results));
  // }

  return (
    <div className="display">
      {survey ? <Survey model={survey} /> : <p>no questionnaire</p>}
    </div>
  );
}

export default App;
