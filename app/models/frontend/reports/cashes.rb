# frozen_string_literal: true

module Frontend
  module Reports
    class Cashes
      def initialize(currency:)
        @currency = currency
      end

      def call
        Cash
          .where(currency:)
          .pluck(:supercategory, :id, :name, :sum)
          .group_by(&:first)
      end

      private

      attr_reader :currency

      class << self
        def call(**kwargs)
          new(**kwargs).call
        end
      end
    end
  end
end
