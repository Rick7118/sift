import { useState } from "react";
import { ExecuteQuery } from "../wailsjs/go/main/App";
import "./App.css";

type QueryResult = {
	Columns: string[];
	Rows: unknown[][];
	RowsAffected: number;
};

function App() {
	const [query, setQuery] = useState("SELECT * FROM users;");
	const [result, setResult] = useState<QueryResult | null>(null);
	const [error, setError] = useState("");
	const [loading, setLoading] = useState(false);

	const runQuery = async () => {
		const sql = query.trim();

		if (!sql) {
			return;
		}

		setLoading(true);
		setError("");

		try {
			const response = await ExecuteQuery(sql);

			setResult(response as QueryResult);
		} catch (err) {
			setResult(null);
			setError(String(err));
		} finally {
			setLoading(false);
		}
	};

	const handleEditorKeyDown = (
		event: React.KeyboardEvent<HTMLTextAreaElement>,
	) => {
		if ((event.ctrlKey || event.metaKey) && event.key === "Enter") {
			event.preventDefault();
			runQuery();
		}
	};

	return (
		<div className="app">
			{/* Sidebar */}
			<aside className="sidebar">
				<div className="sidebar-header">
					<div className="brand">Sift</div>
				</div>

				<div className="sidebar-section">
					<div className="section-header">
						<span>CONNECTIONS</span>
						<button className="icon-button">+</button>
					</div>

					<div className="connection active">
						<span className="status-dot" />

						<div className="connection-info">
							<span>Local SQLite</span>
							<span className="connection-subtitle">
								sift.db
							</span>
						</div>
					</div>
				</div>

				<div className="sidebar-section database-section">
					<div className="section-header">
						<span>DATABASE</span>

						<button className="icon-button">↻</button>
					</div>

					<div className="tree">
						<div className="tree-item folder">
							<span className="chevron">▾</span>
							<span>Tables</span>
						</div>

						<div className="tree-item table">
							<span className="table-icon">▦</span>
							<span>users</span>
						</div>

						<div className="tree-item muted">
							<span className="chevron">›</span>
							<span>Views</span>
						</div>

						<div className="tree-item muted">
							<span className="chevron">›</span>
							<span>Indexes</span>
						</div>
					</div>
				</div>

				<div className="sidebar-footer">
					<button className="sidebar-action">
						<span>?</span>
						<span>Help</span>
					</button>

					<button className="sidebar-action">
						<span>⚙</span>
						<span>Settings</span>
					</button>
				</div>
			</aside>

			{/* Main workspace */}
			<main className="workspace">
				<header className="topbar">
					<div className="tabs">
						<div className="tab active">
							<span className="tab-icon">SQL</span>
							<span>Query 1</span>

							<button className="tab-close">×</button>
						</div>

						<button className="new-tab">+</button>
					</div>

					<div className="topbar-actions">
						<button className="toolbar-button">↻</button>
						<button className="toolbar-button">⚙</button>
					</div>
				</header>

				{/* SQL editor */}
				<section className="editor-panel">
					<div className="panel-toolbar">
						<div className="panel-title">
							<span className="panel-label">SQL</span>

							<span className="connection-label">
								Local SQLite
							</span>
						</div>

						<div className="editor-actions">
							<button
								className="secondary-button"
								onClick={() =>
									setQuery((current) => current.trim())
								}
							>
								Format
							</button>

							<button
								className="run-button"
								onClick={runQuery}
								disabled={loading}
							>
								<span>{loading ? "…" : "▶"}</span>

								{loading ? "Running" : "Run"}

								<span className="shortcut">
									Ctrl + Enter
								</span>
							</button>
						</div>
					</div>

					<div className="editor-container">
						<div className="line-numbers">
							{query.split("\n").map((_, index) => (
								<span key={index}>{index + 1}</span>
							))}
						</div>

						<textarea
							className="sql-editor"
							value={query}
							onChange={(event) =>
								setQuery(event.target.value)
							}
							onKeyDown={handleEditorKeyDown}
							spellCheck={false}
						/>
					</div>
				</section>

				{/* Results */}
				<section className="results-panel">
					<div className="results-toolbar">
						<div className="results-title">
							<span>Results</span>

							{result && result.Columns.length > 0 && (
								<span className="result-count">
									{result.Rows.length}{" "}
									{result.Rows.length === 1
										? "row"
										: "rows"}
								</span>
							)}
						</div>

						<div className="results-actions">
							<button
								className="secondary-button"
								disabled={!result}
							>
								Copy
							</button>

							<button
								className="secondary-button"
								disabled={!result}
							>
								Export
							</button>
						</div>
					</div>

					<div className="results-container">
						{error ? (
							<div className="error-container">
								<div className="error-title">
									Query failed
								</div>

								<div className="error-message">
									{error}
								</div>
							</div>
						) : result && result.Columns.length > 0 ? (
							<div className="table-wrapper">
								<table className="results-table">
									<thead>
										<tr>
											{result.Columns.map((column) => (
												<th key={column}>
													{column}
												</th>
											))}
										</tr>
									</thead>

									<tbody>
										{result.Rows.map((row, rowIndex) => (
											<tr key={rowIndex}>
												{row.map((value, columnIndex) => (
													<td
														key={`${rowIndex}-${columnIndex}`}
													>
														{value === null ||
														value === undefined ? (
															<span className="null-value">
																NULL
															</span>
														) : (
															String(value)
														)}
													</td>
												))}
											</tr>
										))}
									</tbody>
								</table>
							</div>
						) : result ? (
							<div className="success-container">
								<div className="success-title">
									Query executed successfully
								</div>

								<div className="success-description">
									{result.RowsAffected}{" "}
									{result.RowsAffected === 1
										? "row"
										: "rows"}{" "}
									affected.
								</div>
							</div>
						) : (
							<div className="empty-results">
								<div className="empty-icon">⌁</div>

								<div className="empty-title">
									No results yet
								</div>

								<div className="empty-description">
									Run a SQL query to see the results here.
								</div>
							</div>
						)}
					</div>
				</section>

				<footer className="statusbar">
					<div className="status-left">
						<span className="status-item">
							<span className="status-dot" />
							Connected
						</span>

						<span className="status-item">SQLite</span>
					</div>

					<div className="status-right">
						<span>{loading ? "Executing..." : "Ready"}</span>
						<span>Sift v0.1</span>
					</div>
				</footer>
			</main>
		</div>
	);
}

export default App;