export interface Element {
  name: string;
  title: string;
  type: string;
}

export interface Questionnaire {
  elements: Element[];
}

export interface State {
  questionnaire: Questionnaire | null;
}
