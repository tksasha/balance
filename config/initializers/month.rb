# frozen_string_literal: true

class Month
  class << self
    alias original parse

    def parse(*)
      original(*)
    rescue ArgumentError
      now
    end
  end
end
