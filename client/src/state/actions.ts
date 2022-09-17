/* eslint-disable no-param-reassign */
import { Context } from ".";

export const onInitializeOvermind = async ({ state, effects }: Context) => {
  state.isLoading = true;
  state.questionnaire = await effects.api.getQuestionnaire();
  state.isLoading = false;
};
