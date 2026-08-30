export namespace database {
	
	export class Column {
	    Name: string;
	    Type: string;
	    NotNull: boolean;
	    PrimaryKey: boolean;
	
	    static createFrom(source: any = {}) {
	        return new Column(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.Name = source["Name"];
	        this.Type = source["Type"];
	        this.NotNull = source["NotNull"];
	        this.PrimaryKey = source["PrimaryKey"];
	    }
	}
	export class QueryResult {
	    Columns: string[];
	    Rows: any[][];
	    RowsAffected: number;
	
	    static createFrom(source: any = {}) {
	        return new QueryResult(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.Columns = source["Columns"];
	        this.Rows = source["Rows"];
	        this.RowsAffected = source["RowsAffected"];
	    }
	}

}

