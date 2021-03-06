# frozen_string_literal: true

RSpec.describe CategoryWidgetDataSearcher do
  let(:params) { { currency: 'usd' } }

  subject { described_class.new params }

  its(:currency) { should eq 'usd' }

  describe '#search' do
    let :categories do
      [
        ['Food', 13, false],
        ['Drinks', 5, false],
        ['Salary', 17, true],
      ]
    end

    let :collection do
      [
        [
          'Видатки', [
            ['Food', 13],
            ['Drinks', 5],
          ]
        ],
        [
          'Надходження',
          [
            ['Salary', 17],
          ]
        ],
      ]
    end

    before do
      #
      # Category
      #   .visible
      #   .where(currency: 'usd')
      #   .order(:income)
      #   .pluck(:name, :id, :income) -> categories
      #
      allow(Category).to receive(:visible) do
        double.tap do |a|
          allow(a).to receive(:where).with(currency: 'usd') do
            double.tap do |b|
              allow(b).to receive(:order).with(:income) do
                double.tap do |c|
                  allow(c).to receive(:pluck).with(:name, :id, :income).and_return(categories)
                end
              end
            end
          end
        end
      end
    end

    subject { described_class.search params }

    it { should eq collection }
  end

  describe '.search' do
    before do
      #
      # described_class.new(currency: 'usd').search
      #
      expect(described_class).to receive(:new).with(currency: 'usd') do
        double.tap do |a|
          expect(a).to receive(:search)
        end
      end
    end

    subject { described_class.search currency: 'usd' }

    it { expect { subject }.not_to raise_error }
  end
end
