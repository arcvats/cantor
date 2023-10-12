# frozen_string_literal: true

module Structure
  module Heap
    # Defines a Binary heap (default is min heap)
    class BinaryHeap
      attr_reader :size

      def initialize(comparable = nil)
        @elements = []
        @comparable = comparable || ->(a, b) { a < b }
        @size = 0
      end

      def empty?
        @size.zero?
      end

      def peek
        return nil if empty?

        @elements[0]
      end

      def swap(idx, idy)
        @elements[idx], @elements[idy] = @elements[idy], @elements[idx]
      end

      def insert(element)
        @size += 1
        @elements << element
        idx = @size - 1
        parent_idx = idx / 2

        while idx.positive? && @comparable.call(@elements[idx], @elements[parent_idx])
          swap(idx, parent_idx)
          idx = parent_idx
          parent_idx = idx / 2
        end
      end
    end
  end
end
