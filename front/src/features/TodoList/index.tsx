import { TodoData } from "./type";
import { TodoListForm } from "./components/TodoListForm";
import { TodoCheckBox } from "./components/TodoCheckBox";

type Props = {
  data: TodoData;
};

export const TodoList = ({ data }: Props) => {
  return (
    <div className="flex flex-col items-center gap-5 h-screen mt-12">
      <div className="w-[500px]">
        <TodoListForm />
        <TodoCheckBox data={data} />
      </div>
    </div>
  );
};
