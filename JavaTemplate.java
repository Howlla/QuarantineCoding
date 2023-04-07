    import java.io.*;
    import java.util.*;
     
    public class B {
     
        long n, c0, c1;
     
        public void solve() {
            n = in.nextInt() - 1;
            c0 = in.nextInt();
            c1 = in.nextInt();
           
            out.println(result);
        }
     
     
        public void run() {
            in = new FastScanner();
            out = new PrintWriter(System.out);
            solve();
            out.close();
        }
     
        FastScanner in;
        PrintWriter out;
     
        class FastScanner {
            BufferedReader br;
            StringTokenizer st;
     
            public FastScanner(String fileName) {
                try {
                    br = new BufferedReader(new FileReader(fileName));
                } catch (FileNotFoundException e) {
                }
            }
     
            public FastScanner() {
                br = new BufferedReader(new InputStreamReader(System.in));
            }
     
            String nextToken() {
                while (st == null || !st.hasMoreElements()) {
                    try {
                        st = new StringTokenizer(br.readLine());
                    } catch (IOException e) {
                    }
                }
                return st.nextToken();
            }
     
            int nextInt() {
                return Integer.parseInt(nextToken());
            }
     
            long nextLong() {
                return Long.parseLong(nextToken());
            }
     
            double nextDouble() {
                return Double.parseDouble(nextToken());
            }
        }
     
        public static void main(String[] args) {
            new B().run();
        }
    }