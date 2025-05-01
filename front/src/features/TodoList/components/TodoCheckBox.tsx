"use client";

import {
  ApolloQueryResult,
  gql,
  OperationVariables,
  useMutation,
} from "@apollo/client";
import { Checkbox } from "@/components/checkbox";
import { GraphQLResponse } from "../type";

type Props = {
  data: GraphQLResponse;
  refetch: (
    variables?: Partial<OperationVariables> | undefined
  ) => Promise<ApolloQueryResult<GraphQLResponse>>;
};

const CHANGE_STATUS = gql(`
  mutation ChangeTaskStatus($data: ChangeStatus!) {
    changeTaskStatus(data: $data) {
      msg
    }
  }
`);

export const TodoCheckBox = ({ data, refetch }: Props) => {
  const [changeTaskStatus, { error }] = useMutation(CHANGE_STATUS);
  return (
    <div className="mt-2">
      {data.getTasks.map((item) => {
        return (
          <div key={item.ID} className="flex items-center gap-2">
            <Checkbox
              id={item.ID}
              checked={item.isChecked}
              onCheckedChange={async (isCheck) => {
                if (typeof isCheck !== "boolean") return;
                await changeTaskStatus({
                  variables: {
                    data: {
                      id: item.ID,
                      isChecked: isCheck,
                    },
                  },
                });
                await refetch();
              }}
              className="cursor-pointer"
              name="checked"
            />
            <label htmlFor={item.task} className="text-lg ml-2">
              <p>{item.task}</p>
            </label>
          </div>
        );
      })}
      {error && <p className="text-red-500">タスクの更新に失敗しました</p>}
    </div>
  );
};
