// import { useCallback } from "react";

import "survey-core/defaultV2.min.css";
import { StylesManager, Model } from "survey-core";
import { Survey } from "survey-react-ui";

import { useAppState } from "./state";

StylesManager.applyTheme("defaultV2");

function App() {
  const { questionnaire, isLoading } = useAppState();
  const surveyJson = questionnaire;
  const survey = surveyJson && new Model(surveyJson);

  console.log("isLoading", isLoading);

  // // eslint-disable-next-line @typescript-eslint/no-explicit-any
  // const results = useCallback((sender: { [key: string]: any }) => {
  //   return JSON.stringify(sender.data);
  // }, []);

  // if (survey) {
  //   survey.onComplete.add(console.log(results));
  // }

  return isLoading ? (
    <p>...loading</p>
  ) : survey ? (
    <div style={{ width: "80vw" }}>
      <Survey model={survey} />
    </div>
  ) : (
    <p>no questionnaire</p>
  );
}

export default App;
