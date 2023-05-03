import { DataQuery, DataSourceJsonData } from '@grafana/data';


export interface MyQuery extends DataQuery {
  queryText?: string;
  constant: number;
  queryType: string;
  
}

export const DEFAULT_QUERY: Partial<MyQuery> = {
  constant: 6.5,
};

/**
 * These are options configured for each DataSource instance
 */
export interface MyDataSourceOptions extends DataSourceJsonData {
  path?: string;
}

/**
 * Value that is used in the backend, but never sent over HTTP to the frontend
 */
export interface MySecureJsonData {
  apiKey?: string;

}


export interface MyDataSourcePlugin {
  query: (query: MyQuery) => Promise<any>;
  // add other methods as needed
}

export interface MyDataSourceOptions extends DataSourceJsonData {
  maxPoints: number;
}

