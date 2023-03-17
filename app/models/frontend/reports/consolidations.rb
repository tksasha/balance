# frozen_string_literal: true

module Frontend
  module Reports
    class Consolidations
      def initialize(currency:, month:)
        self.currency = currency
        self.month    = month
      end

      def call
        ActiveRecord::Base
          .connection
          .exec_query(sql, 'SQL', bindings)
          .rows
          .group_by(&:first)
          .sort
      end

      private

      attr_reader :currency, :month

      def currency=(currency)
        @currency = Item.currencies[currency] || Item.currencies.values.first
      end

      def month=(month)
        @month = \
          case month
          when String
            Month.parse(month)
          when Month
            month
          end
      end

      def sql
        <<~SQL.squish
          SELECT
            IIF(categories.income, 0, categories.supercategory) AS supercategory,
            categories.name AS category_name,
            categories.id AS category_id,
            ROUND(SUM(sum), 2) AS sum
          FROM
            items
          INNER JOIN categories ON categories.id=items.category_id
          WHERE
            items.currency=$1
          AND
            date BETWEEN $2 AND $3
          AND
            items.deleted_at IS NULL
          GROUP BY
            supercategory, category_name, category_id
        SQL
      end

      def bindings
        [
          ActiveRecord::Relation::QueryAttribute.new('currency', currency, ActiveRecord::Type::String.new),
          ActiveRecord::Relation::QueryAttribute.new('start_date', month.start_date, ActiveRecord::Type::Date.new),
          ActiveRecord::Relation::QueryAttribute.new('end_date', month.end_date, ActiveRecord::Type::Date.new)
        ]
      end

      class << self
        def call(**kwargs)
          new(**kwargs).call
        end
      end
    end
  end
end
