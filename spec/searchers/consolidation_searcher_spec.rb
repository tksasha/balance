# frozen_string_literal: true

RSpec.describe ConsolidationSearcher do
  subject { described_class.new relation, params }

  let(:relation) { double }

  let(:params) { { currency: 'usd', month: '2019-11' } }

  its(:currency) { is_expected.to eq 'usd' }

  its(:month) { is_expected.to eq Month.new(2019, 11) }

  describe '#dates' do
    context do
      let(:dates) { Date.new(2019, 11, 1)..Date.new(2019, 11, 30) }

      its(:dates) { is_expected.to eq dates }
    end

    context do
      let(:dates) { double }

      before { subject.instance_variable_set :@dates, dates }

      its(:dates) { is_expected.to eq dates }
    end
  end

  xdescribe '#search' do
    let(:dates) { Date.new(2019, 11, 1)..Date.new(2019, 11, 30) }

    let(:collection) do
      [
        double(sum: 0.42, income?: false),
        double(sum: 9.00, income?: false),
        double(sum: 2.00, income?: true)
      ]
    end

    before { expect(ConsolidationExpensesSum).to have_received(:sum=).with(9.42) }

    before do
      #
      # relation.
      #   where(date: dates, currency: 'usd').
      #   select('SUM(sum) AS sum, category_id').
      #   group(:category_id) -> collection
      #
      allow(relation).to receive(:where).with(date: dates, currency: 'usd') do
        double.tap do |a|
          allow(a).to receive(:select).with('SUM(sum) AS sum, category_id') do
            double.tap do |b|
              allow(b).to receive(:group).with(:category_id).and_return(collection)
            end
          end
        end
      end
    end

    its(:search) { is_expected.to eq collection }
  end

  xdescribe '.search' do
    subject { described_class.search(:relation, { date: :date }) }

    before do
      #
      # described_class.new(:relation, date: :date).search
      #
      expect(described_class).to have_received(:new).with(:relation, { date: :date }) do
        double.tap { |a| expect(a).to have_received(:search) }
      end
    end

    it { expect { subject }.not_to raise_error }
  end
end
