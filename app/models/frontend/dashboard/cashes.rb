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

      def all
        scope.order(:name)
      end

      def summary
        scope
          .group(:supercategory)
          .sum(:sum)
      end

      private

      attr_reader :currency

      def scope
        ::Cash.where(currency:)
      end

      class << self
        def all(*args)
          new(*args).all
        end
      end
    end
  end
end
