import SearchBar from '../components/SearchBar';

export default function Home() {
  return (
    <main className="home">
      <div className="wordmark" aria-hidden="true">
        AIWIKI
      </div>
      <h1>AIWIKI</h1>
      <p>
        An independent, AI-generated encyclopedia demo. Search any topic and AIWIKI will draft a
        structured article with sections, internal links, generated notes, and an infobox when useful.
      </p>
      <SearchBar size="large" />
      <p className="home-note">Articles are generated locally through Ollama and cached for 10 minutes.</p>
    </main>
  );
}
