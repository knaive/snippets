class Program
{
    static void Main(string[] args)
    {
        const int numTasks = 9;

        // Set up the cancellation source and get the token.
        CancellationTokenSource tokenSource = new CancellationTokenSource();
        CancellationToken token = tokenSource.Token;

        // Set up the tasks
        Task[] tasks = new Task[numTasks];
        for (int i = 0; i < numTasks; i++)
            tasks[i] = Task.Factory.StartNew(() => PerformTask(token), token);

        // Now the tasks are all set up, show the state.
        // Most will be WaitingToRun, some will be Running
        foreach (Task t in tasks.OrderBy(t => t.Id))
            Console.WriteLine("Tasks {0} state: {1}", t.Id, t.Status);

        // Give some of the tasks a chance to do something.
        Thread.Sleep(1500);

        // Cancel the tasks
        Console.WriteLine("Cancelling tasks");
        tokenSource.Cancel();
        Console.WriteLine("Cancellation Signalled");

        try
        {
            // Wait for the tasks to cancel if they've not already completed
            Task.WaitAll(tasks);
        }
        catch (AggregateException aex)
        {
            aex.Handle(ex =>
            {
                // Handle the cancelled tasks
                TaskCanceledException tcex = ex as TaskCanceledException;
                if (tcex != null)
                {
                    Console.WriteLine("Handling cancellation of task {0}", tcex.Task.Id);
                    return true;
                }

                // Not handling any other types of exception.
                return false;
            });
        }

        // Show the state of each of the tasks.
        // Some will be RanToCompletion, others will be Cancelled.
        foreach(Task t in tasks.OrderBy(t => t.Id))
            Console.WriteLine("Tasks {0} state: {1}", t.Id, t.Status);


        Console.WriteLine("Program End");
        Console.ReadLine();
    }

    static void PerformTask(CancellationToken token)
    {
        try
        {
            // The loop simulates work that can be cooperatively cancelled.
            Console.WriteLine("Task {0}: Starting", Task.CurrentId);
            for (int i = 0; i < 4; i++)
            {
                // Check for the cancellation to be signalled
                token.ThrowIfCancellationRequested();

                // Write out a little bit showing the progress of the task
                Console.WriteLine("Task {0}: {1}/4 In progress", Task.CurrentId, i + 1);
                Thread.Sleep(500); // Simulate doing some work
            }
            // By getting here the task will RunToCompletion even if the
            // token has been signalled.
            Console.WriteLine("Task {0}: Finished", Task.CurrentId);
        }
        catch (OperationCanceledException)
        {
            // Any clean up code goes here.
            Console.WriteLine("Task {0}: Cancelling", Task.CurrentId);
            throw; // To ensure that the calling code knows the task was cancelled.
        }
        catch(Exception)
        {
            // Clean up other stuff
            throw; // If the calling code also needs to know.
        }
    }
}