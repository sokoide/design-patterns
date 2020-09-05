namespace patterns
{

    public class Format
    {
        public string FontName;
        public int FontSize;
    }
    public class SpreadSheetCell
    {
        public Format Format;
        public string Data;
    }

    public class SpreadSheet
    {
        private int Width;
        private int Height;

        private SpreadSheetCell[] cells;

        public SpreadSheet(int width, int height)
        {
            Width = width;
            Height = height;
            cells = new SpreadSheetCell[width * height];
            for (int x = 0; x < width; x++)
            {
                for (int y = 0; y < height; y++)
                {
                    cells[width * y + x] = new SpreadSheetCell();
                }
            }
        }

        public void SetFormat(int row, int column, Format format)
        {
            cells[Width * row + column].Format = format;
        }

        public void SetData(int row, int column, string data)
        {
            cells[Width * row + column].Data = data;
        }

        public SpreadSheetCell Cell(int row, int column)
        {
            return cells[Width * row + column];
        }
    }
}