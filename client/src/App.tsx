// import { useCallback } from "react";

import "survey-core/modern.min.css";
import { StylesManager, Model } from "survey-core";
import { Survey } from "survey-react-ui";

import { useAppState } from "./state";

StylesManager.applyTheme("modern");

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

  return (
    <div>
      {isLoading ? (
        <p>...loading</p>
      ) : survey ? (
        <Survey model={survey} />
      ) : (
        <p>no questionnaire</p>
      )}
    </div>
  );
}

export default App;
