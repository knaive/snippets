 var cts = new CancellationTokenSource();

            // start a task that sends a cancel after 500 ms 
            new Task(() =>
            {
                Thread.Sleep(500);
                cts.Cancel();
            }).Start();

            var factory = new TaskFactory(cts.Token);
            Task t1 = factory.StartNew(new Action<object>(f =>
                {
                    Console.WriteLine("in task");
                    for (int i = 0; i < 20; i++)
                    {
                        Thread.Sleep(100);
                        CancellationToken ct = (f as TaskFactory).CancellationToken;
                        if (ct.IsCancellationRequested)
                        {
                            Console.WriteLine("cancelling was requested");
                            ct.ThrowIfCancellationRequested();
                            break;
                        }
                        Console.WriteLine("in loop {0}", i);
                    }
                    Console.WriteLine("task finished");
                }), factory);

            try
            {
                t1.Wait();
            }
            catch (AggregateException ex)
            {
                foreach (var innerEx in ex.InnerExceptions)
                {
                    Console.WriteLine(innerEx.Message);
                }
            }