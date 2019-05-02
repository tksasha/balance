require 'rails_helper'

RSpec.describe ConsolidationSearcher do
  let(:relation) { double }

  let(:params) { {} }

  subject { described_class.new relation, params }

  describe '#date_range' do
    let(:date) { Date.new 2019, 5, 1 }

    let(:params) { { date: date } }

    its(:date_range) { should eq Date.new(2019, 5, 1)..Date.new(2019, 5, 31) }
  end

  describe '#search' do
    let(:date_range) { double }

    let(:collection) do
      [
        double(sum: 0.42, income?: false),
        double(sum: 9.00, income?: false),
        double(sum: 2.00, income?: true)
      ]
    end

    before { expect(subject).to receive(:date_range).and_return(date_range) }

    before { expect(ConsolidationExpensesSum).to receive(:sum=).with(9.42) }

    it do
      #
      # relation.
      #   where(date: date_range).
      #   select('SUM(sum) AS sum, category_id').
      #   group(:category_id) -> collection
      #
      expect(relation).to receive(:where).with(date: date_range) do
        double.tap do |a|
          expect(a).to receive(:select).with('SUM(sum) AS sum, category_id') do
            double.tap do |b|
              expect(b).to receive(:group).with(:category_id).and_return(collection)
            end
          end
        end
      end
    end

    after { subject.search }
  end

  describe '.search' do
    it do
      #
      # described_class.new(:relation, date: :date).search
      #
      expect(described_class).to receive(:new).with(:relation, date: :date) do
        double.tap { |a| expect(a).to receive(:search) }
      end
    end

    after { described_class.search :relation, date: :date }
  end
end
