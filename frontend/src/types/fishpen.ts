export enum FishPenCategory {
  fixed = "fixed",
  floating = "floating",
  submersible = "submersible",
  submersed = "submersed",
  other = "other",
}

export type FishPenResponse = {
  id: number;
  name: string;
  makerModelName?: string;
  description?: string;
  material?: string;
  netMaterial?: string;
  category: FishPenCategory;
  widthCm: number;
  lengthCm: number;
  heightCm: number;
};

export type FishPenRequest = Omit<FishPenResponse, "id">;
