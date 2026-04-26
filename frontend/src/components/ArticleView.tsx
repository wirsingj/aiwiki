import { RefreshCw } from 'lucide-react';
import { Link } from 'react-router-dom';
import { Article, slugify } from '../api';
import Infobox from './Infobox';

type ArticleViewProps = {
  article: Article;
  onRegenerate: () => void;
};

export default function ArticleView({ article, onRegenerate }: ArticleViewProps) {
  const generated = new Date(article.generatedAt).toLocaleString();

  return (
    <main className="article-shell">
      <div className="notice">
        AI-generated article. May contain inaccuracies. Generated notes are not verified citations.
        Verify important claims.
      </div>

      <div className="article-actions">
        {article.cached && <span className="badge">Cached</span>}
        <span className="timestamp">Page last generated {generated}</span>
        <button className="secondary-button" onClick={onRegenerate}>
          <RefreshCw size={15} aria-hidden="true" />
          Regenerate article
        </button>
      </div>

      <article className="article">
        <Infobox infobox={article.infobox} />
        <h1>{article.title}</h1>
        <p className="lead">{article.summary}</p>

        {article.sections.length > 0 && (
          <nav className="toc" aria-label="Contents">
            <h2>Contents</h2>
            <ol>
              {article.sections.map((section, index) => (
                <li key={`${section.heading}-${index}`}>
                  <a href={`#${sectionId(section.heading)}`}>{section.heading}</a>
                </li>
              ))}
            </ol>
          </nav>
        )}

        {article.sections.map((section, index) => (
          <section key={`${section.heading}-${index}`} id={sectionId(section.heading)}>
            <h2>{section.heading}</h2>
            {section.paragraphs.map((paragraph, paragraphIndex) => (
              <p key={`${section.heading}-${paragraphIndex}`}>
                {paragraph}
                {paragraphIndex === section.paragraphs.length - 1 && article.references.length > 0 && (
                  <sup>
                    <a href="#references">[{Math.min(index + 1, article.references.length)}]</a>
                  </sup>
                )}
              </p>
            ))}
            {section.links.length > 0 && (
              <p className="section-links">
                {section.links.map((link) => (
                  <Link key={link} to={`/wiki/${slugify(link)}`}>
                    {link}
                  </Link>
                ))}
              </p>
            )}
          </section>
        ))}

        {article.seeAlso.length > 0 && (
          <section>
            <h2>See also</h2>
            <ul className="see-also">
              {article.seeAlso.map((topic) => (
                <li key={topic}>
                  <Link to={`/wiki/${slugify(topic)}`}>{topic}</Link>
                </li>
              ))}
            </ul>
          </section>
        )}

        <section id="references">
          <h2>References</h2>
          <ol className="references">
            {article.references.map((reference, index) => (
              <li key={`${reference.label}-${index}`}>
                <strong>{reference.label}.</strong> {reference.note}
              </li>
            ))}
          </ol>
        </section>
      </article>
    </main>
  );
}

function sectionId(heading: string) {
  return slugify(heading);
}
