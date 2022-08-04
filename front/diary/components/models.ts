export type listType = { WritedMonths: string[]; Lines: lineType[] };
export type lineType = { Day: string; Outline: string, Tags: string[], IsDetail: boolean, HCount: number };
export type detailType = { Day: string; Outline: string, Tags: string[], Detail: string };
