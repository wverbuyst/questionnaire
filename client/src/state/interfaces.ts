export interface Element {
  name: string;
  title: string;
  type: string;
}

export interface Questionnaire {
  elements: Element[];
}

export interface State {
  isLoading: boolean;
  questionnaire: Questionnaire | null;
}
