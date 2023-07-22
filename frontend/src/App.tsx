import React from "react";
import { Navigate, Route, Routes } from "react-router-dom";

import { FarmHome } from "./FarmHome";
import { FarmList } from "./FarmList";
import { FishPenCreate } from "./FishPenCreate";
import { FishPenEdit } from "./FishPenEdit";

const App: React.FC = () => {
  return (
    <Routes>
      <Route path="/" element={<FarmList />} />
      <Route path="/:farmId" element={<FarmHome />} />
      <Route path="/:farmId/fishpens/:fishPenId" element={<FishPenEdit />} />
      <Route path="/:farmId/fishpens/create" element={<FishPenCreate />} />
      <Route path="*" element={<Navigate to="/" replace />} />
    </Routes>
  );
};

export default App;
