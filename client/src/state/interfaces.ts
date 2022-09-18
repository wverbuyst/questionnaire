export interface Element {
  name: string
  title: string
  type: string
  html?: string
  isRequired?: boolean
  choices?: string[] | null
  correctAnswer?: string
  choicesOrder?: string
}

export interface Questionnaire {
  title: string
  pages: { elements: Element[] }[]
}

export interface State {
  isLoading: boolean
  questionnaire: Questionnaire | null
}
