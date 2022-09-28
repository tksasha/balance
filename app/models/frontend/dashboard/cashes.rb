# frozen_string_literal: true

module Frontend
  class Dashboard
    class Cashes
      def initialize(currency:)
        @currency = currency
      end

      def sum
        scope.sum(:sum)
      end

      private

      attr_reader :currency

      def scope
        Cash.where(currency:)
      end
    end
  end
end
