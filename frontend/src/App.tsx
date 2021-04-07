import { useQuery } from "@urql/preact";
import { h } from "preact";

const query = `
query GetAllUsers {
  allUsers {
    id
    username
    dateCreated
    messagesCount
  }
}
`;

const FirstPage = () => {
  const {} = useQuery({ query });

  return <div>page 1</div>;
};

export const App = () => {
  return <FirstPage />;
};
