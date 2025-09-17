import { buildModule } from "@nomicfoundation/hardhat-ignition/modules";

const HashRegistryModule = buildModule("HashRegistryModule", (m) => {
  const hashRegistry = m.contract("HashRegistry");

  return { hashRegistry };
});

export default HashRegistryModule;