import { createClient } from "@supabase/supabase-js";

const supabaseUrl = "https://wthkdylfqxzaknqcnvzp.supabase.co"
const supabaseAnonPublic = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpc3MiOiJzdXBhYmFzZSIsInJlZiI6Ind0aGtkeWxmcXh6YWtucWNudnpwIiwicm9sZSI6ImFub24iLCJpYXQiOjE3MDYwNDU3NTUsImV4cCI6MjAyMTYyMTc1NX0.nY5D2dmGbQcURS0stir9HsQBH19M_B8vhrwkrGEcQyI"

export const supabase = createClient(supabaseUrl, supabaseAnonPublic);