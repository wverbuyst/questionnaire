// import { useCallback } from "react";

import 'survey-core/defaultV2.min.css'

import { Model, StylesManager } from 'survey-core'
import { Survey } from 'survey-react-ui'

import { useAppState } from './state'

StylesManager.applyTheme('defaultV2')

function App() {
  const { countries, questionnaire, isLoading } = useAppState()
  const surveyJson = questionnaire
  const survey = surveyJson && new Model(surveyJson)

  survey?.setVariable('countries', countries)

  return isLoading ? (
    <p>...loading</p>
  ) : survey ? (
    <div style={{ width: '80vw' }}>
      <Survey model={survey} />
    </div>
  ) : (
    <p>no questionnaire</p>
  )
}

export default App
