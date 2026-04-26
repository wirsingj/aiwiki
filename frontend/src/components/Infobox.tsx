import { Infobox as InfoboxType } from '../api';

type InfoboxProps = {
  infobox?: InfoboxType;
};

export default function Infobox({ infobox }: InfoboxProps) {
  if (!infobox || infobox.rows.length === 0) return null;

  return (
    <aside className="infobox" aria-label={`${infobox.heading} facts`}>
      <h2>{infobox.heading}</h2>
      <dl>
        {infobox.rows.map((row) => (
          <div key={`${row.label}-${row.value}`}>
            <dt>{row.label}</dt>
            <dd>{row.value}</dd>
          </div>
        ))}
      </dl>
    </aside>
  );
}
