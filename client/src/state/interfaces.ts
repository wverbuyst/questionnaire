export interface Element {
  name: string;
  title: string;
  type: string;
}

export interface Elements {
  element: Element[];
}

export interface Questionnaire {
  title: string;
  pages: Elements;
}

export interface State {
  isLoading: boolean;
  questionnaire: Questionnaire | null;
}
