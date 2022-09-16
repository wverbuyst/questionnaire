/* eslint-disable no-param-reassign */
import { Context } from ".";

export const onInitializeOvermind = async ({ state, effects }: Context) => {
  state.questionnaire = await effects.api.getQuestionnaire();
};
