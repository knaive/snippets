using System;
using System.Collections;
using System.Collections.Generic;
using System.Web.Script.Serialization;
using System.Management.Automation;


namespace CmdletTest
{
    [Cmdlet(VerbsLifecycle.Start, "CmdTest")]
    public class Program : PSCmdlet
    {
        [Parameter(ValueFromPipelineByPropertyName = true)]
        public string Collect { get; set; }

        protected override void ProcessRecord()
        {
            var serializer = new JavaScriptSerializer();
            Dictionary<string, object> collectPara = null;
            collectPara = serializer.Deserialize<Dictionary<string, object>>(Collect);
            foreach (KeyValuePair<string, object> item in collectPara)
            {
                Console.WriteLine(item.Key + ": " + item.Value);
            }
        }
    }
}
