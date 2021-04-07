export type Maybe<T> = T | null;
export type Exact<T extends { [key: string]: unknown }> = { [K in keyof T]: T[K] };
export type MakeOptional<T, K extends keyof T> = Omit<T, K> & { [SubKey in K]?: Maybe<T[SubKey]> };
export type MakeMaybe<T, K extends keyof T> = Omit<T, K> & { [SubKey in K]: Maybe<T[SubKey]> };
/** All built-in and custom scalars, mapped to their actual values */
export type Scalars = {
  ID: string;
  String: string;
  Boolean: boolean;
  Int: number;
  Float: number;
  Time: any;
};

export type Chat = {
  __typename?: 'Chat';
  id: Scalars['ID'];
  title: Scalars['String'];
  members: Array<User>;
  membersCount: Scalars['Int'];
  messages: Array<Message>;
  messageCount: Scalars['Int'];
  dateCreated: Scalars['Time'];
  dateUpdated: Scalars['Time'];
};

export type CreateChatPayload = {
  title: Scalars['String'];
};

export type CreateUserInput = {
  username: Scalars['String'];
  password: Scalars['String'];
};

export type JoinChatPayload = {
  userId: Scalars['String'];
  chatId: Scalars['String'];
};

export type Message = {
  __typename?: 'Message';
  id: Scalars['ID'];
  text: Scalars['String'];
  senderId: Scalars['String'];
  sender: User;
  chatId: Scalars['String'];
  chat: Chat;
  dateCreated: Scalars['Time'];
};

export type Mutation = {
  __typename?: 'Mutation';
  createUser: User;
  sendMessage: Message;
  createChat: Chat;
  joinChat: Chat;
};


export type MutationCreateUserArgs = {
  input?: Maybe<CreateUserInput>;
};


export type MutationSendMessageArgs = {
  input?: Maybe<SendMessageInput>;
};


export type MutationCreateChatArgs = {
  input?: Maybe<CreateChatPayload>;
};


export type MutationJoinChatArgs = {
  input?: Maybe<JoinChatPayload>;
};

export type Query = {
  __typename?: 'Query';
  allUsers: Array<User>;
  allMessages: Array<Message>;
  allChats: Array<Chat>;
};

export type SendMessageInput = {
  text: Scalars['String'];
  senderId: Scalars['String'];
  chatId: Scalars['String'];
};


export type User = {
  __typename?: 'User';
  id: Scalars['ID'];
  username: Scalars['String'];
  messagesCount: Scalars['Int'];
  dateCreated: Scalars['Time'];
  dateUpdated: Scalars['Time'];
};
