export type TodoData = {
  id: string;
  task: string;
  checked: boolean;
}[];

export type GraphQLResponse = {
  getTasks: {
    ID: string;
    task: string;
    isChecked: boolean;
  }[];
};
