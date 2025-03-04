"use client";

import { useMemo } from "react";
import { DEFAULT_DRAGGABLE_WIDTH } from "../../../constants";
import { useLogsContext } from "../../../context/logs";
import { extractResponseField, safeParseJson } from "../../../utils";
import { LogFooter } from "./components/log-footer";
import { LogHeader } from "./components/log-header";
import { LogMetaSection } from "./components/log-meta";
import { LogSection } from "./components/log-section";
import { ResizablePanel } from "./resizable-panel";

const PANEL_MAX_WIDTH = 600;
const PANEL_MIN_WIDTH = 400;

const createPanelStyle = (distanceToTop: number) => ({
  top: `${distanceToTop}px`,
  width: `${DEFAULT_DRAGGABLE_WIDTH}px`,
  height: `calc(100vh - ${distanceToTop}px)`,
  paddingBottom: "1rem",
});

type Props = {
  distanceToTop: number;
};

export const LogDetails = ({ distanceToTop }: Props) => {
  const { setSelectedLog, selectedLog: log } = useLogsContext();
  const panelStyle = useMemo(() => createPanelStyle(distanceToTop), [distanceToTop]);

  if (!log) {
    return null;
  }

  const handleClose = () => {
    setSelectedLog(null);
  };

  return (
    <ResizablePanel
      minW={PANEL_MIN_WIDTH}
      maxW={PANEL_MAX_WIDTH}
      onClose={handleClose}
      className="absolute right-0 bg-gray-1 dark:bg-black font-mono drop-shadow-2xl overflow-y-auto z-20 p-4"
      style={panelStyle}
    >
      <LogHeader log={log} onClose={handleClose} />

      <LogSection details={log.request_headers} title="Request Header" />
      <LogSection
        details={JSON.stringify(safeParseJson(log.request_body), null, 2)}
        title="Request Body"
      />
      <LogSection details={log.response_headers} title="Response Header" />
      <LogSection
        details={JSON.stringify(safeParseJson(log.response_body), null, 2)}
        title="Response Body"
      />
      <div className="mt-3" />
      <LogFooter log={log} />
      <LogMetaSection content={JSON.stringify(extractResponseField(log, "meta"), null, 2)} />
    </ResizablePanel>
  );
};
