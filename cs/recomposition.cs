 public class MEFRecompExportPerson1
  {
    [Export]
    public int Age { get { return 20; } }
  }

  public class MEFRecompExportPerson2
  {
    [Export]
    public int Age { get { return 25; } }
  }

  public class MEFRecompImportPerson
  {
    [ImportMany(AllowRecomposition = true)]
    public int[] AgeCollection { get; set; }
  }

  class MEFRecompTest
  {
    public static void Main(string[] args)
    {
      AggregateCatalog catalog = new AggregateCatalog(new TypeCatalog(typeof(MEFRecompExportPerson1)));
      CompositionContainer compositionContainer = new CompositionContainer(catalog);

      MEFRecompImportPerson importInstance = new MEFRecompImportPerson();

      compositionContainer.ComposeParts(importInstance);
      DisplayAge("Initial compose with 1 item in the catalog MEFRecompExportPerson1", importInstance.AgeCollection);
      //result : 
      //20

      catalog.Catalogs.Add(new TypeCatalog(typeof(MEFRecompExportPerson2)));
      DisplayAge("Composition after adding MEFRecompExportPerson2 to the catalog", importInstance.AgeCollection);
      //result : 
      //20
      //25

      Console.ReadLine();
    }

    private static void DisplayAge(string message, int[] ageCollection)
    {
      Console.WriteLine(message + " : ");
      foreach (var ageItem in ageCollection)
      {        
        Console.WriteLine(ageItem.ToString());        
      }
      Console.WriteLine("-------------------");
    }
  }